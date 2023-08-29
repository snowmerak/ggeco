import axios from "axios";

/*
* {
  "badges": [
    {
      "id": "string",
      "image": "string",
      "name": "string",
      "summary": "string"
    }
  ]
}
* */
export class Badge {
    id: string;
    name: string;
    summary: string;
    active_image: string;
    inactive_image: string;
    selected_image: string;
}

export async function get_badges(url: string, key: string): Promise<Badge[]> {
    url = `${url}/badge/list`;

    let resp = await axios.get(url, {
        headers: {
            'x-functions-key': key
        }
    });

    if (resp.status != 200) {
        throw new Error("Error getting badges");
    }

    if (!resp.data.badges) {
        throw new Error("Success But no badges");
    }

    return resp.data.badges as Badge[];
}