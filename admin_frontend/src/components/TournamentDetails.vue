<template src="../templates/tournament_details.html">
    
</template>

<script>
    import EventBus from './event-bus';
    import UpdatePrizeForm from './UpdatePrizeForm.vue';
    export default {
        name: "TournamentDetails",
        components:{
            UpdatePrizeForm
        },
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
                players: [],
                elements: [],
                elementsCount:0,
                winners:[],
                prizes:[],
                prizesShow:false
            }

        },
        methods:{
            getPlayers(id){
                this.elements = [];
                this.prizes = [];
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
            },
            addWinner(){
               this.elements.push({type: 'UpdatePrizeForm',id:this.elementsCount+=1});
            },
            savePrize(winner){
                this.winners.push(winner);
                for (let i = 0; i < this.winners.length; i++) {
                   this.winners[i].tournamentId = this.tournamentId;
                }
                console.log(this.winners);
            },
            updatePrizes(){
                this.elements = [];
                this.$http.put(this.$hostname+'/update_prizes', this.winners)
                    .then(response => {
                        if (this.statusCode!==200){
                            this.errorMsg = response.data;
                            this.errorAlert = true;
                        }
                        this.winners = [];
                        this.$http.get(this.$hostname+'/result_tournament/'+this.tournamentId)
                            .then(response => {
                                if (this.statusCode!==200){
                                    this.errorMsg = response.data;
                                    this.errorAlert = true;
                                }
                                this.prizes = response.data.msg.winners;
                                this.prizesShow = true;

                            }, error => {
                                this.errorMsg = error.statusText || error.data;
                                this.statusCode = error.statusCode || error.status;
                                this.errorAlert = true;
                            });

                    }, error => {
                        this.errorMsg = error.statusText || error.data;
                        this.statusCode = error.statusCode || error.status;
                        this.errorAlert = true;
                    });
            },
            clearWinners(){
                this.winners = [];
            }

        },
        created: function () {
            EventBus.$on('init',this.getPlayers);

        },
        mounted: function () {
            EventBus.$on('prize',this.savePrize);
            EventBus.$on('hide',this.clearWinners);
        }


    }
</script>

<style scoped src="../styles/tournament_details.css">

</style>