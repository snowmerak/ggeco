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
        TextInput, Toggle
    } from "carbon-components-svelte";
    import {add_badge} from "./add_badge";
    import {config} from "../config/config";

    let badges = [];

    let name = "";
    let summary = "";
    let activeImageFiles = [];
    let inactiveImageFiles = [];
    let selectedImageFiles = [];
    let searchable = false;
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
        { key: "active_image", value: "활성 이미지" },
        { key: "inactive_image", value: "비활성 이미지" },
        { key: "selected_image", value: "선택 이미지" },
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
            <br/>
            <TextInput
                labelText="이름"
                placeholder="배지 이름을 입력해주세요."
                bind:value={name}
            ></TextInput>
            <br/>
            <TextInput
                labelText="설명"
                placeholder="배지 설명을 입력해주세요."
                bind:value={summary}
            ></TextInput>
            <br/>
            <FileUploader
                    bind:files={activeImageFiles}
                    labelDescription="활성 배지 이미지를 선택해주세요."
                    buttonLabel="파일 선택"
                    accept={[".png", ".jpg", ".jpeg", ".gif", ".webp"]}
                    status="edit"
            ></FileUploader>
            <br />
            <FileUploader
                    bind:files={inactiveImageFiles}
                    labelDescription="비활성 배지 이미지를 선택해주세요."
                    buttonLabel="파일 선택"
                    accept={[".png", ".jpg", ".jpeg", ".gif", ".webp"]}
                    status="edit"
            ></FileUploader>
            <br />
            <FileUploader
                    bind:files={selectedImageFiles}
                    labelDescription="선택 배지 이미지를 선택해주세요."
                    buttonLabel="파일 선택"
                    accept={[".png", ".jpg", ".jpeg", ".gif", ".webp"]}
                    status="edit"
            ></FileUploader>
            <br/>
            <Toggle
                labelText="검색 가능"
                bind:toggled={searchable}
            ></Toggle>
            <br/>
            <Button kind="secondary" on:click={async () => {
                alert("파일 업로드 중입니다. 잠시만 기다려주세요.");
                try {
                    let id = await add_badge($config.back_url, $config.back_key, name, summary, activeImageFiles[0], inactiveImageFiles[0], selectedImageFiles[0], searchable);
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
