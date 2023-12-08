<script>
    import tabler from 'yesvelte/css/tabler.min.css?url';
    import Login from './lib/components/Login.svelte';
    import Pag1 from './lib/components/Pag1.svelte';
    import Pag2 from './lib/components/Pag2.svelte';
    import Settings from './lib/components/Settings.svelte';
    import {SidebarItem, Sidebar, El, ButtonGroup,Button, Icon} from "yesvelte";

    let isLogin = false
    let selectedMenu = "login";


</script>
<svelte:head>
    <link rel='stylesheet' href={tabler}/>
</svelte:head>

<main data-theme="dark" data-bs-theme="dark">
    <El container style="overflow: hidden">
        <El row>
            <El col="2">
                <Sidebar theme="dark" style="width: fit-content">
                    {#if isLogin}
                    <ButtonGroup col="auto">
                        <Button>
                            <Icon name="settings" on:click={() => (selectedMenu = "settings")}/>
                        </Button>
                        <Button>
                            <Icon name="user" />
                        </Button>
                    </ButtonGroup>
                    {/if}
                    {#if !isLogin}
                        <SidebarItem icon="login" title="Identificarse" on:click={() => (selectedMenu = "login")}/>
                    {:else}
                        <SidebarItem icon="file" title="Page 1" on:click={() => (selectedMenu = "pag1")}/>
                        <SidebarItem icon="file" title="Page 2" on:click={() => (selectedMenu = "pag2")}/>
                    {/if}
                </Sidebar>
            </El>
            <El col mt="5">
                {#if isLogin}
                    {#if selectedMenu === "pag1"}
                        <Pag1 />
                    {:else if selectedMenu === "pag2"}
                        <Pag2 />
                    {:else if selectedMenu === "settings"}
                        <Settings />
                    {:else}
                        <h1>
                            Page Not Found
                        </h1>
                    {/if}
                {:else}
                    <Login bind:loggedIn={isLogin} bind:redirectAfterLogin={selectedMenu}/>
                {/if}
            </El>
        </El>
    </El>
</main>

