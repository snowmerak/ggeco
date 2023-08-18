import {writable} from "svelte/store";

export let config = writable({
    "app_url": "",
    "app_key": "",
    "back_url": "",
    "back_key": "",
});

class Config {
    app_url: string;
    app_key: string;
    back_url: string;
    back_key: string;
}

export function set_config(data: string) {
    let value = JSON.parse(data) as Config;
    config.set({
        "app_url": value?.app_url ?? "",
        "app_key": value?.app_key ?? "",
        "back_url": value?.back_url ?? "",
        "back_key": value?.back_key ?? "",
    });
    localStorage.setItem("config", data);
}