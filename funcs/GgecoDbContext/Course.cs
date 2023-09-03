using System;
using System.Collections.Generic;

namespace funcs;

public partial class Course
{
    public Guid Id { get; set; }

    public Guid AuthorId { get; set; }

    public string Name { get; set; }

    public DateTime RegDate { get; set; }

    public string Review { get; set; }

    public byte IsPublic { get; set; }
}
