<script>
    import tabler from 'yesvelte/css/tabler.min.css?url';
    import {El} from "yesvelte";

    import Login from './lib/components/Login.svelte';
    import Pag1 from './lib/components/Pag1.svelte';
    import Pag2 from './lib/components/Pag2.svelte';
    import Settings from './lib/components/Settings.svelte';
    import SideNavBar from './lib/components/SideNavBar.svelte';
    import TopNavBar from './lib/components/TopNavBar.svelte';

    let isLogin = false
    let selectedMenu = "login";


</script>
<svelte:head>
    <link rel='stylesheet' href={tabler}/>
</svelte:head>

<main data-theme="dark" data-bs-theme="dark">
    <TopNavBar />
    <SideNavBar {isLogin} {selectedMenu} on:menuChange={(e) => (selectedMenu = e.detail)}/>

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
</main>

<style>
    main {
        height: 100vh;
        width: 100vw;
        overflow: hidden;
    }
</style>

