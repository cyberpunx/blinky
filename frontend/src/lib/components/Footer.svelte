<script>
import {El} from "yesvelte";
</script>

<footer class="site-footer">
    <El row>
        <El col="12" colSm="4" textAlign="start" px="3" textMuted style="background: var(--ds-color3)">Hogwarts Rol️</El>
        <El col="12" colSm="4" textAlign="center" px="3" textMuted style="background: var(--ds-color3)">️</El>
        <El col="12" colSm="4" textAlign="end" px="3" textMuted style="background: var(--ds-color3)">Author: Aiden Ward - Made with ❤️</El>
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