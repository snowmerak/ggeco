<script lang="ts">
    import "carbon-components-svelte/css/g10.css";
  import {Tab, TabContent, Tabs} from "carbon-components-svelte";
    import BadgeList from "./badge/Badges.svelte";
    import Config from "./config/Config.svelte";

  let selected = 0;

  let tabLabels = [
      "Config",
      "Badge",
  ];
</script>

<Tabs bind:selected>
    {#each tabLabels as label}
        <Tab>{label}</Tab>
    {/each}
  <svelte:fragment slot="content">
      <TabContent>
          <Config />
      </TabContent>
      <TabContent>
          <BadgeList />
      </TabContent>
  </svelte:fragment>
</Tabs>