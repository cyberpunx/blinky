<script>
    import { El, Navbar, NavbarItem, Avatar, Dropdown, DropdownItem, DropdownMenu, Button, Icon } from 'yesvelte'
    import {MENU} from '../constants.js';
    import {createEventDispatcher} from "svelte";

    export let isLogin;
    export let selectedMenu;
    export let username ="";
    export let initials = "";

    const dispatch = createEventDispatcher();

    function handleMenuChange(menu) {
        dispatch('menuChange', menu);
    }
</script>

{#if isLogin}
<El row alignItems="center" justifyContent="between" style="margin-left: 0; margin-right: 0">
    <Navbar col="12" colSm="1" style="width: fit-content;">
        <NavbarItem icon="menu-2">
            <NavbarItem title="Pag1" on:click={() => handleMenuChange(MENU.PAG1)} />
            <NavbarItem title="Pag2" on:click={() => handleMenuChange(MENU.PAG2)} />
            <NavbarItem title="SETTINGS" on:click={() => handleMenuChange(MENU.SETTINGS)} />
        </NavbarItem>
    </Navbar>

    <El col="12" colSm="4" tag="h1" fontWeight="bolder" style="color:var(--ds-color3)"> -= HR HELPER =-</El>

    <El h="25" col="12" colSm="2" style="width: fit-content;">
        <Dropdown  arrow={false}>
            <Avatar color="primary" size="md" shape="rounded">{initials}</Avatar>
            <DropdownMenu>
                <DropdownItem disabled>{username}</DropdownItem>
                <DropdownItem>Item 1</DropdownItem>
                <DropdownItem>Item 2</DropdownItem>
            </DropdownMenu>
        </Dropdown>
    </El>
</El>
{/if}