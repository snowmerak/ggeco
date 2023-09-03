using System;
using System.Collections.Generic;

namespace funcs;

public partial class User
{
    public Guid Id { get; set; }

    public string Nickname { get; set; }

    public byte? Age { get; set; }

    public byte? Gender { get; set; }

    public DateTime CreateAt { get; set; }

    public DateTime LastSignin { get; set; }

    public Guid Badge { get; set; }

    public DateTime? RemovedAt { get; set; }
}
