<template src="../templates/tournament_details.html">
    
</template>

<script>
    import EventBus from './event-bus';
    export default {
        name: "TournamentDetails",
        data(){
            return {
                tournamentId: null,
                statusCode: null,
                errorMsg: '',
                errorAlert: false,
                fields: {
                   id: {
                        label: 'Id'
                    },
                    name:{
                        label: 'Name'
                    },
                    points:  {
                        label: 'Balance',
                        sortable: true
                    } },
                perPage: 5,
                currentPage: 1,
                totalRows:0,
                players: []
            }

        },
        methods:{
            getPlayers(id){
                this.tournamentId = id;
                this.$http.get(this.$hostname+'/get_players_tournament/'+this.tournamentId)
                    .then(response => {
                        this.players = response.data.msg;
                        this.totalRows = this.players.length;
                        this.statusCode = response.data.code;
                        if (this.statusCode!==200){
                            this.errorMsg = response.data;
                            this.errorAlert = true;
                        }
                    }, error => {
                        this.errorMsg = error.statusText || error.data;
                        this.statusCode = error.statusCode || error.status;
                        this.errorAlert = true;
                    });
            }


        },
        created: function () {
            EventBus.$on('init',this.getPlayers)
        }


    }
</script>

<style scoped src="../styles/tournament_details.css">

</style>