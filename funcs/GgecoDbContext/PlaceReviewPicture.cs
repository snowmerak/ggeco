using System;
using System.Collections.Generic;

namespace funcs;

public partial class PlaceReviewPicture
{
    public Guid Id { get; set; }

    public Guid ReviewId { get; set; }

    public int Order { get; set; }

    public string PictureUrl { get; set; }

    public string ThumbnailUrl { get; set; }
}
