<script>
    import {Card, CardBody, El, FormAutocomplete,FormInput, Input, Button, Autocomplete } from 'yesvelte'
    import {GetPotionSubforum, UpdatePotionSubforum, GetPotionThread, UpdatePotionThread} from "../../../wailsjs/go/main/App.js";
    import {onMount} from "svelte";
    export let config;
    export let tool;
    let baseUrl = config["baseUrl"]

    onMount(() => {
        GetPotionSubforum().then((result) => {
            result.forEach((item) => {
                valuePotionSub = [...valuePotionSub, item.url];
                itemsPotionSub = [...itemsPotionSub, item.url];
            })
            timeLimitPotionSub = result[0].timeLimit;
            turnLimitPotionSub = result[0].turnLimit;
        })

        GetPotionThread().then((result) => {
            result.forEach((item) => {
                valuePotionThr = [...valuePotionThr, item.url];
                itemsPotionThr = [...itemsPotionThr, item.url];
            })
            timeLimitPotionThr = result[0].timeLimit;
            turnLimitPotionThr = result[0].turnLimit;
        })
    });

    let itemsPotionSub = [];
    let valuePotionSub = [];
    let timeLimitPotionSub;
    let turnLimitPotionSub;

    let itemsPotionThr = [];
    let valuePotionThr = [];
    let timeLimitPotionThr;
    let turnLimitPotionThr;

    function onCreatedPotionSub({ detail }) {
        if (detail.startsWith(baseUrl)) {
            detail = detail.replace(baseUrl, "");
        }
        detail = detail.split("#")[0];
        detail = detail.split("?")[0];
        console.log("onCreatedPotionSub");
        valuePotionSub = [...valuePotionSub, detail];
        itemsPotionSub = [...itemsPotionSub, detail];
    }

    function onCreatedPotionThr({ detail }) {
        if (detail.startsWith(baseUrl)) {
            detail = detail.replace(baseUrl, "");
        }
        detail = detail.split("#")[0];
        detail = detail.split("?")[0];
        console.log("onCreatedPotionThr");
        valuePotionThr = [...valuePotionThr, detail];
        itemsPotionThr = [...itemsPotionThr, detail];
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
        })
    }

</script>

<El>
    <El tag="h1" textAlign="start">Configuración</El>
    <El tag="hr" />
    <El row>
        <El col="12" colSm="6">
            <Card title="Pociones [Subforos]">
                <CardBody textAlign="start">
                    <Autocomplete dismissible on:created={onCreatedPotionSub} multiple create items={itemsPotionSub}
                                  bind:value={valuePotionSub}/>
                    <FormInput mt="3" label="Tiempo Límite" type="number" placeholder="72" bind:value={timeLimitPotionSub} />
                    <FormInput mt="3" label="Cantidad de Turnos" type="number" placeholder="8" bind:value={turnLimitPotionSub}/>
                    <Button mt="3" color="success" on:click={doUpdatePotionSubforum}>Guardar cambios</Button>
                </CardBody>
            </Card>
            <El tag="hr" />
        </El>
        <El col="12" colSm="6">
            <Card title="Pociones [Temas]">
                <CardBody textAlign="start">
                    <Autocomplete dismissible on:created={onCreatedPotionThr} multiple create items={itemsPotionThr}
                                  bind:value={valuePotionThr}/>
                    <FormInput mt="3" label="Tiempo Límite" type="number" placeholder="72" bind:value={timeLimitPotionThr} />
                    <FormInput mt="3" label="Cantidad de Turnos" type="number" placeholder="8" bind:value={turnLimitPotionThr}/>
                    <Button mt="3" color="success" on:click={doUpdatePotionThread}>Guardar cambios</Button>
                </CardBody>
            </Card>
            <El tag="hr" />
        </El>
    </El>

</El>
