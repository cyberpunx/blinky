<script>
    import {SidebarItem, Sidebar, ButtonGroup, Button, Icon} from "yesvelte";
    import { createEventDispatcher } from 'svelte';

    export let isLogin;
    export let selectedMenu;

    const dispatch = createEventDispatcher();

    function handleMenuChange(menu) {
        dispatch('menuChange', menu);
    }
</script>

<Sidebar theme="dark" style="width: fit-content; position: relative;">
    {#if isLogin}
        <ButtonGroup col="auto">
            <Button>
                <Icon name="settings" on:click={() => handleMenuChange("settings")}/>
            </Button>
            <Button>
                <Icon name="user" />
            </Button>
        </ButtonGroup>
    {/if}
    {#if !isLogin}
        <SidebarItem icon="login" title="Identificarse" on:click={() => handleMenuChange("login")}/>
    {:else}
        <SidebarItem icon="file" title="Page 1" on:click={() => handleMenuChange("pag1")}/>
        <SidebarItem icon="file" title="Page 2" on:click={() => handleMenuChange("pag2")}/>
    {/if}
</Sidebar>