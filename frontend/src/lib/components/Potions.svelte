<script>
    import {Card,CardBody,CardTitle,CardFooter,CardActions, El, Spinner, Table, TableBody, TableCell, TableHead, TableRow,
        Accordions, Accordion, AccordionBody, Badge, Icon} from 'yesvelte'
    import {onMount} from "svelte";
    import {GetPotionSubforum, GetPotionThread, SubforumPotionsClub} from "../../../wailsjs/go/main/App.js";
    import {HOUSES} from '../constants.js';
    import {getUserColorByHouse} from '../util.js';
    import {BrowserOpenURL} from "../../../wailsjs/runtime/runtime.js"

    export let config;
    export let tool;
    let loading = true;
    let subForumPotionsThreads;


    onMount(() => {
        let potionThrUrls = [];
        let timeLimitPotionThr;
        let turnLimitPotionThr;

        GetPotionSubforum().then((result) => {
            let potionSubUrls = [];
            let timeLimitPotionSub = result[0].timeLimit;
            let turnLimitPotionSub = result[0].turnLimit;

            result.forEach((item) => {
                potionSubUrls = [...potionSubUrls, item.url];
            })

            SubforumPotionsClub(potionSubUrls, timeLimitPotionSub, turnLimitPotionSub).then((result) => {
                console.log(result['threadReports']);
                subForumPotionsThreads = result['threadReports'];
                loading = false;
            })
        })



        GetPotionThread().then((result) => {
            console.log(result);
            if (result !== null) {
                result.forEach((item) => {
                    potionThrUrls = [...potionThrUrls, item.url];
                })
                timeLimitPotionThr = result[0].timeLimit;
                turnLimitPotionThr = result[0].turnLimit;
            }
        })
    });

    function buildThreadTitle(thread) {
        const player1 = thread['Player1']['Name'];
        const player2 = thread['Player2']['Name'];
        const potion = thread['Potion']['Name'];
        const moderator = thread['Moderator']['Name'];
        const p1house = thread['Player1']['House'];
        const p2house = thread['Player2']['House'];
        const p1color = getUserColorByHouse(p1house);
        const p2color = getUserColorByHouse(p2house);

        const threadNames = `<span style="color: var(${p1color})">${player1}</span> & <span style="color: var(${p2color})">${player2}</span>`;
        const potionAndModerator = `(${potion}) <span>Mod: <span style="color: var(--mod-color)">${moderator}</span></span>`;

        return `${threadNames} ${potionAndModerator}`;
    }

    function buildPlayerName(player) {
        const name = player['Name'];
        const house = player['House'];
        const color = getUserColorByHouse(house);

        return `<span style="color: var(${color})">${name}</span>`;
    }

    function getPlayerUrl(player) {
        return config['baseUrl'] + player['ProfileUrl'];
    }

    function getPlayerHouseColor(player) {
        const house = player['House'];
        return getUserColorByHouse(house);
    }

    function getStatusColor(status) {
        switch (status) {
            case 'Success':
                return 'green';
            case 'Fail':
                return 'red';
            case 'WaitingPlayer1':
                return 'light';
            case 'WaitingPlayer2':
                return 'light';
            default:
                return 'yellow';
        }
    }

    function getStatusText(thread) {
        const status = thread['Status'];
        const player1 = thread['Player1']['Name'];
        const player2 = thread['Player2']['Name'];
        const p1color = getUserColorByHouse(thread['Player1']['House']);
        const p2color = getUserColorByHouse(thread['Player2']['House']);
        const p1url = thread['Player1']['Url'];
        const p2url = thread['Player2']['Url'];

        switch (status) {
            case 'Success':
                return 'Éxito';
            case 'Fail':
                return 'Fracaso';
            case 'WaitingPlayer1':
                if (player1 !=="") {
                    return `Esperando a <span style="color: var(${p1color})">${player1}</span>`;
                }else{
                    return `Esperando a Jugador 1`;
                }
            case 'WaitingPlayer2':
                if (player2 !=="") {
                    return `Esperando a <span style="color: var(${p2color})">${player2}</span>`;
                }else{
                    return `Esperando a Jugador 2`;
                }
            default:
                return 'Desconocido';
        }
    }

    function getPotionName(thread) {
        const potion = thread['Potion']['Name'];
        return potion;
    }

    function getPotionUrl(thread) {
        const url = config['baseUrl'] + thread['Thread']['Url'];
        return url;
    }

    function getModeratorName(thread) {
        const moderator = thread['Moderator']['Name'];
        return moderator;
    }

    function getModeratorUrl(thread) {
        const url = config['baseUrl'] + thread['Thread']['Author']['Url'];
        return url;
    }

    function getElapsedTime(thread) {
        //elapsedTime := forumDateTime.Sub(*post.Created)
        const forumTime = tool['ForumDateTime'] //format is "2023-12-18T21:35:00Z"
        const postList = thread['Thread']['Posts'];
        const lastPost = postList[postList.length - 1];
        const postTime = lastPost['Created']; //format is "2023-12-18T21:35:00Z"
        const forumDateTime = new Date(forumTime);
        const postDateTime = new Date(postTime);
        const elapsedTime = Math.abs(forumDateTime - postDateTime);

        let h,m,s;
        h = Math.floor(elapsedTime/1000/60/60);
        m = Math.floor((elapsedTime/1000/60/60 - h)*60);
        return `${h}h ${m}m`;

    }



</script>

<El textAlign="start">
    <El row>
        <El col="12" colSm="1">
            <El tag="h1">Pociones</El>
        </El>
        <El col="12" colSm="6" pt="2" >
            <El tag="a" style="cursor: pointer; color:--mod-color" on:click={BrowserOpenURL('https://docs.google.com/spreadsheets/d/' + config["gSheetId"])}>
                Hoja de Moderación <Icon name="external-link" />
            </El>
        </El>
    </El>

    {#if loading}
        <El textAlign="center">
            <Spinner size="lg" color="primary" />
            <p>Cargando pociones...</p>
        </El>
    {:else}
        {#each subForumPotionsThreads as thread}
            <El col="12" colSm="12" mt="5">
                <Card status statusColor="{getStatusColor(thread['Status'])}" statusPosition="start" statusSize="md">
                    <CardBody>
                        <CardTitle>
                            <El row>
                                <El col="12" colSm="8">
                                    [ <El tag="a" style="cursor: pointer; color:--mod-color" on:click={BrowserOpenURL(getModeratorUrl(thread))}>
                                        {@html getModeratorName(thread)}
                                    </El>]
                                    <El tag="a" style="cursor: pointer" on:click={BrowserOpenURL(getPlayerUrl(thread['Player1']))}>
                                        {@html buildPlayerName(thread["Player1"])}</El> &
                                    <El tag="a" style="cursor: pointer" on:click={BrowserOpenURL(getPlayerUrl(thread['Player2']))}>
                                        {@html buildPlayerName(thread["Player2"])}</El>
                                    <El tag="a" style="cursor: pointer; color:white" on:click={BrowserOpenURL(getPotionUrl(thread))}>
                                        ({@html getPotionName(thread)})
                                    </El>
                                    <El tag="span" textColor="muted" fontSize="5">| Tiempo transcurrido: {getElapsedTime(thread)} |</El>
                                </El>
                                <El col="12" colSm="4" textAlign="end" textColor="{getStatusColor(thread['Status'])}">
                                    {@html getStatusText(thread)}
                                </El>
                            </El>
                        </CardTitle>
                        <p class="text-muted">
                            <Accordions>
                                <Accordion title="Ver detalle de turnos">
                                    <AccordionBody>
                                        <Table border>
                                            <TableHead>
                                                <TableRow>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">Turno</TableCell>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">Jugador</TableCell>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">Dado</TableCell>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">A tiempo</TableCell>
                                                </TableRow>
                                            </TableHead>
                                            <TableBody>
                                                {#each thread['Turns'] as turn}
                                                <TableRow>
                                                    <TableCell>{turn['Number']}</TableCell>
                                                    <TableCell>{turn['Player']['Role']} | {turn['Player']['Name']}</TableCell>
                                                    <TableCell>{turn['DiceValue']}</TableCell>
                                                    <TableCell>
                                                        {#if turn['OnTime']}
                                                            <Icon name="check" color="success"/>
                                                        {:else}
                                                            {#if turn['DayOffUsed']}
                                                                <Icon name="check" color="warning"/>
                                                            {:else}
                                                                <Icon name="x" color="danger"/>
                                                            {/if}
                                                        {/if}
                                                    </TableCell>
                                                </TableRow>
                                                {/each}
                                            </TableBody>
                                        </Table>
                                    </AccordionBody>
                                </Accordion>
                            </Accordions>
                        </p>
                    </CardBody>
                    <CardFooter>This is standard card footer</CardFooter>
                </Card>
            </El>
            <El tag="hr" />
        {/each}
    {/if}

</El>