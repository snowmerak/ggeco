using System;
using System.Collections.Generic;

namespace funcs;

public partial class BadgeRank
{
    public Guid UserId { get; set; }

    public long CurrentRank { get; set; }

    public long? PrevRank { get; set; }

    public DateTime? UpdateAt { get; set; }
}
