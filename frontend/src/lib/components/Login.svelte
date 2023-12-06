<script>
    import {Login} from '../../../wailsjs/go/main/App'
    import tabler from 'yesvelte/css/tabler.min.css?url'
    import {
        El,
        Button,
        Card,
        CardBody,
        Input,
        FormField,
        Label,
        Alert
    } from 'yesvelte'

    let user, pass
    let error = false
    export let loggedIn = false
    export const redirectAfterLogin = 1
    function doLogin(){
        Login(user, pass).then((result) => {
            if(result){
                loggedIn = true
            }else{
                loggedIn = false
                error = true
            }
        })
    }

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
                            Credenciales inválidas.
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
                    <Button mt="3" col="12" color="primary" on:click={doLogin}>Ingresar</Button>
                </CardBody>
            </Card>

        </El>
        <El col></El>
    </El>
    <El row></El>
</El>


Texto