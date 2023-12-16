<script>
    import {Card,CardBody,CardTitle,CardFooter,CardActions, El, Spinner, Table, TableBody, TableCell, TableHead, TableRow, Accordions, Accordion, AccordionBody} from 'yesvelte'
    import {onMount} from "svelte";
    import {GetPotionSubforum, GetPotionThread, SubforumPotionsClub} from "../../../wailsjs/go/main/App.js";
    import {HOUSES} from '../constants.js';
    import {getUserColorByHouse} from '../util.js';

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

    function getStatusColor(status) {
        switch (status) {
            case 'Success':
                return 'green';
            case 'Fail':
                return 'red';
            case 'WaitingPlayer1':
                return 'dark';
            case 'WaitingPlayer2':
                return 'dark';
            default:
                return 'yellow';
        }
    }

</script>

<El textAlign="start">
    <h1>Pociones</h1>

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
                            {@html buildThreadTitle(thread)}
                        </CardTitle>
                        <p class="text-muted">
                            <Accordions>
                                <Accordion title="Ver detalle de turnos">
                                    <AccordionBody>
                                        <Table border>
                                            <TableHead>
                                                <TableRow>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">Turno</TableCell>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">Rol</TableCell>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">Dado</TableCell>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">Jugador</TableCell>
                                                    <TableCell style="color: var(--ds-color8); background-color: var(--ds-color2)">A tiempo</TableCell>
                                                </TableRow>
                                            </TableHead>
                                            <TableBody>
                                                {#each thread['Turns'] as turn}
                                                <TableRow>
                                                    <TableCell>{turn['Number']}</TableCell>
                                                    <TableCell>{turn['Player']['Role']}</TableCell>
                                                    <TableCell>{turn['Player']['Name']}</TableCell>
                                                    <TableCell>{turn['DiceValue']}</TableCell>
                                                    <TableCell>{turn['OnTime']}</TableCell>
                                                </TableRow>
                                                {/each}
                                            </TableBody>
                                        </Table>
                                    </AccordionBody>
                                </Accordion>
                            </Accordions>
                    </CardBody>
                </Card>
            </El>
            <El tag="hr" />
        {/each}
    {/if}

</El>