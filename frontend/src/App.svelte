<script>
    import tabler from 'yesvelte/css/tabler.min.css?url';
    import {El} from "yesvelte";

    import {MENU} from './lib/constants';
    import Login from './lib/components/Login.svelte';
    import Home from './lib/components/Home.svelte';
    import Potions from './lib/components/Potions.svelte';
    import Settings from './lib/components/Settings.svelte';
    import TopNavBar from './lib/components/TopNavBar.svelte';
    import Footer from './lib/components/Footer.svelte';

    let isLogin = false
    let selectedMenu = "login";
    let username = "";
    let initials = "";
    let config
    let tool


</script>
<svelte:head>
    <link rel='stylesheet' href={tabler}/>
</svelte:head>

<main data-theme="dark" data-bs-theme="dark">
    <TopNavBar {initials} {username} {isLogin} {selectedMenu} on:menuChange={(e) => (selectedMenu = e.detail)}/>

    <El p="3">
    {#if isLogin}
        {#if selectedMenu === MENU.HOME}
            <Home />
        {:else if selectedMenu === MENU.POTIONS}
            <Potions {config} {tool} />
        {:else if selectedMenu === MENU.SETTINGS}
            <Settings {config} {tool} />
        {:else}
            <h1>
                Page Not Found
            </h1>
        {/if}
    {:else}
        <Login bind:tool={tool} bind:config={config} bind:loggedIn={isLogin}
               bind:redirectAfterLogin={selectedMenu} bind:username={username} bind:initials={initials}/>
    {/if}
    </El>

    <Footer />
</main>

<style>
    main {
        min-height: 100%;
        margin-bottom: -20px;
    }
    main:after {
        content: "";
        display: block;
    }

</style>

