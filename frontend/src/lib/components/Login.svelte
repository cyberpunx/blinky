<script>
    import {MENU} from '../constants.js';
    import {Login} from '../../../wailsjs/go/main/App'
    import {GetConfig} from '../../../wailsjs/go/main/App'
    import tabler from 'yesvelte/css/tabler.min.css?url'
    import { onMount } from 'svelte';
    import {
        El,
        Button,
        Card,
        CardBody,
        Input,
        FormField,
        Label,
        Alert,
        Checkbox
    } from 'yesvelte'

    let user, pass
    let error = false
    let message = ""
    let rememberMe = false
    let config = {}
    export let loggedIn = false
    export const redirectAfterLogin = MENU.PAG1
    export let username = ""
    export let initials = ""
    function doLogin(){
        Login(user, pass, rememberMe).then((result) => {
            console.log(result)
            loggedIn = result["success"]
            message = result["message"]
            username = result["username"]
            initials = result["initials"]
            if(!loggedIn) {
                error = true
            }
        })
    }

    onMount(() => {
        GetConfig().then((result) => {
            config = result
            console.log(config)
            if (result["remember"]){
                rememberMe = true
                user = result["username"]
                pass = result["password"]
            }
        })
    });

</script>
<svelte:head>
    <link rel='stylesheet' href={tabler}/>
</svelte:head>


<El container style="overflow: hidden">
    <El row style="height:20vh;"></El>
    <El row>
        <El col></El>
        <El col>

            <Card size="md">
                <CardBody>
                    {#if error}
                        <Alert important icon="alert-circle" color="danger">
                            {message}
                        </Alert>
                    {/if}
                    <El tag="h1">Iniciar Sesión</El>
                    <FormField>
                        <El d="flex" justifyContent="between">
                            <Label>Usuario</Label>
                        </El>
                        <Input placeholder="Tu Usuario" bind:value={user}></Input>
                    </FormField>
                    <FormField mt="3">
                        <El d="flex" justifyContent="between">
                            <Label>Contraseña</Label>
                        </El>
                        <Input type="password" placeholder="Tu Contraseña" bind:value={pass}>
                        </Input>
                    </FormField>
                    <FormField mt="3">
                        <Checkbox bind:checked={rememberMe}>Recordarme
                        </Checkbox>
                    </FormField>
                    <Button mt="3" col="12" color="primary" on:click={doLogin}>Ingresar</Button>
                </CardBody>
            </Card>

        </El>
        <El col></El>
    </El>
    <El row></El>
</El>
