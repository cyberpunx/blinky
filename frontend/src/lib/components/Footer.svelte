<script>
import {El} from "yesvelte";
import {onMount} from "svelte";
import {GetConfig} from '../../../wailsjs/go/main/App'
import {GetTool} from '../../../wailsjs/go/main/App'

export let tool

let formattedTime

function startClock(initialDateTime) {
    // Parse the initial date and time from the string
    let currentTime = new Date(initialDateTime);

    // Check if the initial dateTime is valid
    if (isNaN(currentTime.getTime())) {
        console.error("Invalid initial dateTime");
        return; // Exit the function if the dateTime is invalid
    }

    // Function to update the clock
    function updateClock() {
        // Add one second
        currentTime = new Date(currentTime.getTime() + 1000);

        // Check if the updated time is valid
        if (isNaN(currentTime.getTime())) {
            console.error("Invalid time value encountered");
            clearInterval(clockInterval); // Stop the interval if invalid time is encountered
            return;
        }

        // Format the time as desired
        // Extract UTC date components
        let day = currentTime.getUTCDate().toString().padStart(2, '0');
        let month = (currentTime.getUTCMonth() + 1).toString().padStart(2, '0'); // Month is 0-indexed
        let year = currentTime.getUTCFullYear();
        let hours = currentTime.getUTCHours().toString().padStart(2, '0');
        let minutes = currentTime.getUTCMinutes().toString().padStart(2, '0');
        let seconds = currentTime.getUTCSeconds().toString().padStart(2, '0');

        // Format the time as 'DD-MM-YYYY HH:MM:SS'
        formattedTime = `${day}/${month}/${year} ${hours}:${minutes}`;
    }

    // Update the clock every second
    let clockInterval = setInterval(updateClock, 1000);
}

onMount(() => {
    console.log(tool["ForumDateTime"])
    startClock(tool["ForumDateTime"]);
});

</script>

<footer class="site-footer">
    <El row>
        <El col="12" colSm="4" textAlign="start" px="3" textMuted style="background: var(--ds-color3)">
            <El tag="span" fontWeight="bold">Hogwarts Rol</El>
        </El>
        <El col="12" colSm="4" textAlign="center" px="3" textMuted style="background: var(--ds-color3)">
            {#if !formattedTime}
                Cargando fecha del foro...
            {:else}
                <El tag="span" fontWeight="bold">[ {formattedTime} ]</El>
            {/if}
        ️</El>
        <El col="12" colSm="4" textAlign="end" px="3" textMuted style="background: var(--ds-color3)">
            <El tag="span" fontWeight="bold">Author:</El>
            <El tag="span" textColor="info">Aiden Ward</El> - Made with ❤️
        </El>
    </El>
</footer>

<style>
    .site-footer, main:after {
        height: 20px;
    }
    .site-footer {
        position: fixed;
        bottom: 0;
        right: 0;
        width: 100%;
        text-align:center;
        background-color:orange;
    }
</style>