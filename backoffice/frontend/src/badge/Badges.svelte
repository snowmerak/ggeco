<script lang="ts">
    import {get_badges} from "./get_badges";
    import {
        Button,
        DataTable,
        FileUploader,
        ImageLoader,
        Tab,
        TabContent,
        Tabs,
        TextInput
    } from "carbon-components-svelte";
    import {add_badge} from "./add_badge";
    import {config} from "../config/config";

    let badges = [];

    let name = "";
    let summary = "";
    let files = [];
</script>

<h2>배지</h2>

<Tabs>
    <Tab>전체 목록</Tab>
    <Tab>추가</Tab>
    <svelte:fragment slot="content">
        <TabContent>
            <Button on:click={async () => {
                try {
                    badges = await get_badges($config.app_url, $config.app_key);
                } catch (e) {
                    alert(`배지 목록을 불러오는데 실패했습니다. ${e}`);
                }
            }}>
                Reload
            </Button>
            <DataTable
                    sortable
                    headers={[
        { key: "id", value: "ID" },
        { key: "name", value: "이름" },
        { key: "summary", value: "설명" },
        { key: "image", value: "이미지" },
    ]}
                    rows={badges}
            >
                <svelte:fragment slot="cell" let:row let:cell>
                    {#if cell.key === "image"}
                        <ImageLoader src={cell.value} alt={row.name}/>
                    {:else}
                        {cell.value}
                    {/if}
                </svelte:fragment>
            </DataTable>
        </TabContent>
        <TabContent>
            <h3>배지 추가</h3>
            <TextInput
                labelText="이름"
                placeholder="배지 이름을 입력해주세요."
                bind:value={name}
            ></TextInput>
            <TextInput
                labelText="설명"
                placeholder="배지 설명을 입력해주세요."
                bind:value={summary}
            ></TextInput>
            <FileUploader
                    bind:files
                    labelTitle="배지 이미지"
                    labelDescription="배지 이미지를 선택해주세요."
                    buttonLabel="파일 선택"
                    accept={[".png", ".jpg", ".jpeg", ".gif", ".webp"]}
                    status="edit"
            ></FileUploader>
            <Button kind="secondary" on:click={async () => {
                alert(`파일: ${files[0].name}을 업로드합니다.`);
                try {
                    let id = await add_badge($config.back_url, $config.back_key, name, summary, files[0]);
                    alert(`배지가 추가되었습니다. ID: ${id}`);
                } catch (e) {
                    alert(`파일 업로드에 실패했습니다. ${e}`);
                    return;
                }
            }}>
                추가
            </Button>
        </TabContent>
    </svelte:fragment>
</Tabs>
