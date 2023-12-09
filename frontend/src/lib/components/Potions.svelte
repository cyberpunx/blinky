<script>
    import { El, Table, TableBody, TableCell, TableHead, TableRow } from 'yesvelte'
</script>

<El textAlign="start">
    <h1>Pociones</h1>

    <Table striped>
        <TableHead >
            <TableRow>
                <TableCell style="color: var(--ds-color1)">Participantes</TableCell>
                <TableCell style="color: var(--ds-color1)">Poción</TableCell>
                <TableCell style="color: var(--ds-color1)">Moderador</TableCell>
                <TableCell style="color: var(--ds-color1)">Role</TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            <TableRow>
                <TableCell>Paweł Kuna</TableCell>
                <TableCell>UI Designer, Training</TableCell>
                <TableCell>paweluna@howstuffworks.com</TableCell>
                <TableCell>User</TableCell>
            </TableRow>

            <TableRow>
                <TableCell>Jeffie Lewzey</TableCell>
                <TableCell>Chemical Engineer, Support</TableCell>
                <TableCell>jlewzey1@seesaa.net</TableCell>
                <TableCell>Admin</TableCell>
            </TableRow>

            <TableRow>
                <TableCell>Mallory Hulme</TableCell>
                <TableCell>Geologist IV, Support</TableCell>
                <TableCell>mhulme2@domainmarket.com</TableCell>
                <TableCell>User</TableCell>
            </TableRow>

            <TableRow>
                <TableCell>Dunn Slane</TableCell>
                <TableCell>Research Nurse, Sales</TableCell>
                <TableCell>dslane3@epa.gov</TableCell>
                <TableCell>Owner</TableCell>
            </TableRow>

            <TableRow>
                <TableCell>Emmy Levet</TableCell>
                <TableCell>VP Product Management, Accounting</TableCell>
                <TableCell>elevet4@senate.gov</TableCell>
                <TableCell>Admin</TableCell>
            </TableRow>
        </TableBody>
    </Table>
</El>