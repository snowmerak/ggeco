using System;
using System.Linq;
using Microsoft.Azure.WebJobs;
using Microsoft.Data.SqlClient;
using Microsoft.Extensions.Logging;

namespace funcs
{
    public class UpdateBadgeRank
    {
        private const string SqlServerDatabase = "SQL_SERVER_DATABASE";
        private const string SqlServerHost = "SQL_SERVER_HOST";
        private const string SqlServerPassword = "SQL_SERVER_PASSWORD";
        private const string SqlServerPort = "SQL_SERVER_PORT";
        private const string SqlServerUser = "SQL_SERVER_USER";
        
        [FunctionName("UpdateBadgeRank")]
        // every monday 03:00 AM
        public void Run([TimerTrigger("0 0 3 * * Mon")]TimerInfo myTimer, ILogger log)
        {
            log.LogInformation($"Update Badge Rank executed at: {DateTime.Now}");

            var connectionStringBuilder = new SqlConnectionStringBuilder()
            {
                DataSource = Environment.GetEnvironmentVariable(SqlServerDatabase) ?? "block",
                InitialCatalog = Environment.GetEnvironmentVariable(SqlServerHost) ?? "localhost" + ":" + Environment.GetEnvironmentVariable(SqlServerPort) ?? "1433",
                Password = Environment.GetEnvironmentVariable(SqlServerPassword) ?? "password",
                UserID = Environment.GetEnvironmentVariable(SqlServerUser) ?? "sa",
                IntegratedSecurity = false,
            };
            
            using var db = new GgecoDbContext(connectionStringBuilder.ToString());
            
            using var transaction = db.Database.BeginTransaction();

            try
            {
                // group by count of user's earned badges count
                var ranks = db.Users
                    .Join(
                        db.EarnedBadges,
                        user => user.Id,
                        earnedBadge => earnedBadge.UserId,
                        (user, earnedBadge) => new { userId = user.Id, badgeId = earnedBadge.BadgeId }
                    ).GroupBy(
                        x => x.userId,
                        x => x.badgeId,
                        (key, g) => new { userId = key, count = g.Count() }
                    ).OrderByDescending(x => x.count).ToList();

                var now = DateTime.Now;

                // update badge rank
                for (int i = 0; i < ranks.Count(); i++)
                {
                    var rank = db.BadgeRanks.FirstOrDefault(badgeRank => badgeRank.UserId == ranks[i].userId);
                    if (rank == null)
                    {
                        rank = new BadgeRank()
                        {
                            UserId = ranks[i].userId,
                            PrevRank = 0,
                            CurrentRank = i + 1,
                            UpdateAt = now
                        };
                        db.BadgeRanks.Add(rank);
                    }
                    else
                    {
                        rank.PrevRank = rank.CurrentRank;
                        rank.CurrentRank = i + 1;
                        rank.UpdateAt = now;
                        db.BadgeRanks.Update(rank);
                    }

                    db.SaveChanges();
                }

                transaction.Commit();
            }
            catch (Exception)
            {
                db.Database.RollbackTransaction();
            }

            log.LogInformation($"Update Badge Rank completed at: {DateTime.Now}");
        }
    }
}
