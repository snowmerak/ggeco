using System;
using System.Collections.Generic;

namespace funcs;

public partial class FavoriteCourse
{
    public Guid Id { get; set; }

    public Guid UserId { get; set; }

    public Guid CourseId { get; set; }

    public DateTime? RegisteredAt { get; set; }
}
