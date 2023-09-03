using System;
using System.Collections.Generic;

namespace funcs;

public partial class UserPlaceVisitCount
{
    public Guid Id { get; set; }

    public Guid UserId { get; set; }

    public string PlaceType { get; set; }

    public long Count { get; set; }
}
