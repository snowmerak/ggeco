import axios from "axios";

export async function add_badge(url: string, key: string, name: string, summary: string, active_image_data: File, inactive_image_data: File, selected_image_data: File, searchable: boolean): Promise<string> {
    url = `${url}/badge`;

    const formData = new FormData();
    formData.append("name", name);
    formData.append("description", summary);
    formData.append("active_image", active_image_data);
    formData.append("inactive_image", inactive_image_data);
    formData.append("selected_image", selected_image_data);
    formData.append("searchable", searchable ? "true" : "false");

    let resp = await axios.post(url, formData, {
        headers: {
            'x-functions-key': key
        }
    });

    if (resp.status != 200) {
        throw new Error("Error adding badge");
    }

    if (!resp.data.id) {
        throw new Error("Success But no id");
    }

    return resp.data.id as string;
}