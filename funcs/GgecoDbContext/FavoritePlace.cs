using System;
using System.Collections.Generic;

namespace funcs;

public partial class FavoritePlace
{
    public Guid Id { get; set; }

    public Guid UserId { get; set; }

    public string PlaceId { get; set; }

    public DateTime? RegisteredAt { get; set; }
}
