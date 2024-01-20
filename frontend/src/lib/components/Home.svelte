<script>
    import {
        El, Card, CardBody,
    } from 'yesvelte'
    import {GetThread} from "../../../wailsjs/go/main/App.js";
    export let config;
    export let tool;

    const baseUrl = config['baseUrl']
    const checkTimeThread = GetThread(baseUrl + 't24-001-normas-del-foro')
    console.log(checkTimeThread)

</script>

<El px="5">
    <El row>
        <El col="12" colSm="4"></El>
        <El col="12" colSm="5">
            <El tag="hr" />
            <Card title="Importante">
                <CardBody>
                <El textAlign="start" tag="p">
                    Antes de continuar, ve al foro y <a href="https://www.hogwartsrol.com/profile?mode=editprofile&page_profil=preferences" target="_blank">configura tu perfil</a> con el siguiente formato de fecha y hora:
                    <El tag="kbd">DD/MM/YYYY, HH:mm</El>.
                    Por ejemplo, el día de nochebuena un minuto antes de navidad es <El tag="kbd">24/12/2023, 23:59</El>
                </El>
                </CardBody>
            </Card>
            <El tag="hr" />
        </El>

    </El>
</El>

<style>
     .mono {
        font-family: monospace;
    }
</style>