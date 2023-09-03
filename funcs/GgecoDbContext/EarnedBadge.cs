using System;
using System.Collections.Generic;

namespace funcs;

public partial class EarnedBadge
{
    public Guid Id { get; set; }

    public Guid UserId { get; set; }

    public Guid BadgeId { get; set; }

    public DateTime EarnedAt { get; set; }
}
