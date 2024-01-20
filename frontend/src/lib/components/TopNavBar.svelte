<script>
    import { El, Navbar, NavbarItem, Avatar, Dropdown, DropdownItem, DropdownMenu, Button, Icon } from 'yesvelte'
    import {MENU} from '../constants.js';
    import {createEventDispatcher} from "svelte";
    import {pickRandomColor} from '../util.js';

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
<El row alignItems="center" justifyContent="between" alignSelf="center" style="margin-left: 0; margin-right: 0; background-color: var(--ds-color3)">
    <Navbar col="12" colSm="1" style="width: fit-content;" py="0" my="0" g="1" >
        <Avatar color={pickRandomColor()} size="sm" shape="rounded" alignSelf="center">{initials}</Avatar>
        <NavbarItem>
            <DropdownItem disabled>{username}</DropdownItem>
            <NavbarItem icon="home" title="Inicio" on:click={() => handleMenuChange(MENU.HOME)} />
            <NavbarItem icon="flask-2-filled" title="Pociones" on:click={() => handleMenuChange(MENU.POTIONS)} />
            <DropdownItem divider />
            <NavbarItem icon="settings" title="Configuración" on:click={() => handleMenuChange(MENU.SETTINGS)} />
        </NavbarItem>
    </Navbar>

    <El col="12" colSm="4" fontWeight="bolder" fontSize="3" style="color:var(--ds-color5)">✨ Dobby! 💫</El>

    <El h="25" col="12" colSm="2" style="width: fit-content;">
    </El>
</El>
{/if}