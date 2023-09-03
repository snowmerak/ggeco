using System;
using System.Collections.Generic;

namespace funcs;

public partial class CoursePlace
{
    public Guid Id { get; set; }

    public Guid CourseId { get; set; }

    public string PlaceId { get; set; }

    public int Order { get; set; }
}
