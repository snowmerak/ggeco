<script lang="ts">
    import "carbon-components-svelte/css/g10.css";
  import {Tab, TabContent, Tabs} from "carbon-components-svelte";

  let selected = 0;

  let tabLabels = [
          "Badge",
  ];

  let tabContents = [
          "Badge",
  ];
</script>

<Tabs bind:selected>
    {#each tabLabels as label}
        <Tab>{label}</Tab>
    {/each}
  <svelte:fragment slot="content">
    {#each tabContents as content}
        <TabContent>{content}</TabContent>
    {/each}
  </svelte:fragment>
</Tabs>