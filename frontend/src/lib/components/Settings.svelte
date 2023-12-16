<script>
    import {Card, CardBody, El, FormAutocomplete,FormInput, Input, Button, Autocomplete, Alert } from 'yesvelte'
    import {GetPotionSubforum, UpdatePotionSubforum, GetPotionThread, UpdatePotionThread, UpdateSheetConfig, GetConfig} from "../../../wailsjs/go/main/App.js";
    import {onMount} from "svelte";
    import {MENU} from "../constants.js";

    export let config;
    export let tool;

    let baseUrl = config["baseUrl"]

    onMount(() => {
        loadData();
    });

    function loadData(){
        GetConfig().then((result) => {
            config = result;
            baseUrl = config["baseUrl"]
            credentialsFile = config["gSheetCredFile"];
            tokenFile = config["gSheetTokenFile"]
            sheetId = config["gSheetId"]
        })

        GetPotionSubforum().then((result) => {
            valuePotionSub = [];
            itemsPotionSub = [];
            result.forEach((item) => {
                valuePotionSub = [...valuePotionSub, item.url];
                itemsPotionSub = [...itemsPotionSub, item.url];
            })
            timeLimitPotionSub = result[0].timeLimit;
            turnLimitPotionSub = result[0].turnLimit;
        })

        GetPotionThread().then((result) => {
            valuePotionThr = [];
            itemsPotionThr = [];
            result.forEach((item) => {
                valuePotionThr = [...valuePotionThr, item.url];
                itemsPotionThr = [...itemsPotionThr, item.url];
            })
            timeLimitPotionThr = result[0].timeLimit;
            turnLimitPotionThr = result[0].turnLimit;
        })
    }
    let saveConfirmation = '';
    let open = false;

    let credentialsFile = config["gSheetCredFile"];
    let tokenFile = config["gSheetTokenFile"]
    let sheetId = config["gSheetId"]

    let itemsPotionSub = [];
    let valuePotionSub = [];
    let timeLimitPotionSub = 72;
    let turnLimitPotionSub = 8;

    let itemsPotionThr = [];
    let valuePotionThr = [];
    let timeLimitPotionThr = 48;
    let turnLimitPotionThr = 4;

    function onCreatedPotionSub({ detail }) {
        if (detail.startsWith(baseUrl)) {
            detail = detail.replace(baseUrl, "");
        }
        detail = detail.split("#")[0];
        detail = detail.split("?")[0];
        if (!valuePotionSub.includes(detail)) {
            valuePotionSub = [...valuePotionSub, detail];
            itemsPotionSub = [...itemsPotionSub, detail];
        }
    }

    function onCreatedPotionThr({ detail }) {
        if (detail.startsWith(baseUrl)) {
            detail = detail.replace(baseUrl, "");
        }
        detail = detail.split("#")[0];
        detail = detail.split("?")[0];
        if (!valuePotionThr.includes(detail)) {
            valuePotionThr = [...valuePotionThr, detail];
            itemsPotionThr = [...itemsPotionThr, detail];
        }
    }

    function doUpdatePotionSubforum() {
        let valuePotionSubMap = valuePotionSub.map((subforumUrl) => {
            return {
                url: subforumUrl,
                timeLimit: Number(timeLimitPotionSub),
                turnLimit: Number(turnLimitPotionSub)
            }
        });

        UpdatePotionSubforum(valuePotionSubMap).then((result) => {
            loadData();
        })

    }

    function doUpdatePotionThread() {
        let valuePotionThrMap = valuePotionThr.map((threadUrl) => {
            return {
                url: threadUrl,
                timeLimit: Number(timeLimitPotionThr),
                turnLimit: Number(turnLimitPotionThr)
            }
        });

        UpdatePotionThread(valuePotionThrMap).then((result) => {
            loadData();
        })
    }

    function doUpdateSheetConfig() {
        UpdateSheetConfig(tokenFile, credentialsFile, sheetId).then((result) => {
            loadData();
        })
    }

    function saveChanges(){
        doUpdateSheetConfig();
        doUpdatePotionSubforum();
        doUpdatePotionThread();
        showSaveConfirmation('Cambios guardados correctamente');
    }

    function showSaveConfirmation(message) {
        saveConfirmation = message;
        open = true;
        setTimeout(() => {
            saveConfirmation = '';
            open = false;
        }, 3000);
    }

</script>

<El>
    <El row>
        <El col="12" colSm="2" textAlign="start"><El tag="h1">Configuración</El></El>
        <El col="12" colSm="8">
            <Alert dismissible important bind:open icon="info-circle" color="info">
                {saveConfirmation}
            </Alert>
        </El>
        <El col="12" colSm="2" textAlign="end"><Button color="success" on:click={saveChanges}>Guardar cambios</Button></El>
    </El>

    <El tag="hr" mt="2" mb="2"/>

    <El row rowCols="3">
        <El col="12" colSm="4">
            <Card title="Google Sheet de Moderación">
                <CardBody textAlign="start">
                    <FormInput mt="3" label="Sheet Id" type="text"  bind:value={sheetId} />
                    <FormInput mt="3" label="Credentials File" type="text"  bind:value={credentialsFile} />
                    <FormInput mt="3" label="Token File" type="text" bind:value={tokenFile} />
                </CardBody>
            </Card>
        </El>
        <El col="12" colSm="4">
            <Card title="Pociones [Subforos]">
                <CardBody textAlign="start">
                    <Autocomplete dismissible on:created={onCreatedPotionSub} multiple create items={itemsPotionSub}
                                  bind:value={valuePotionSub}/>
                    <FormInput mt="3" label="Tiempo Límite" type="number" bind:value={timeLimitPotionSub} />
                    <FormInput mt="3" label="Cantidad de Turnos" type="number"  bind:value={turnLimitPotionSub}/>
                </CardBody>
            </Card>
        </El>
        <El col="12" colSm="4">
            <Card title="Pociones [Temas]">
                <CardBody textAlign="start">
                    <Autocomplete dismissible on:created={onCreatedPotionThr} multiple create items={itemsPotionThr}
                                  bind:value={valuePotionThr}/>
                    <FormInput mt="3" label="Tiempo Límite" type="number" bind:value={timeLimitPotionThr} />
                    <FormInput mt="3" label="Cantidad de Turnos" type="number" bind:value={turnLimitPotionThr}/>
                </CardBody>
            </Card>
        </El>
    </El>


</El>
