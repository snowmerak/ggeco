<script lang="ts">
    import {Button, TextArea} from "carbon-components-svelte";
    import {set_config} from "./config";

    let data = localStorage.getItem('settings');
</script>

<TextArea
    bind:value={data}
    labelText="Settings"
    placeholder="input setting here"
></TextArea>
<Button
    on:click={() => {
        set_config(data)
    }}
>
    Save
</Button>