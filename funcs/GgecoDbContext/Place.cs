using System;
using System.Collections.Generic;

namespace funcs;

public partial class Place
{
    public string Id { get; set; }

    public byte[] Data { get; set; }

    public DateTime? LastUpdate { get; set; }
}
