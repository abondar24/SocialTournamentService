<template src="../templates/tournament.html">
    
</template>

<script>
    import AnnounceForm from './AnnounceForm.vue'
    import TournamentDetails from './TournamentDetails'
    import EventBus from './event-bus';
    export default {
        name: "Tournament",
        components:{
            AnnounceForm,
            TournamentDetails,
        },

        data() {
            return {

                tournaments: [],
                statusCode: null,
                errorMsg: '',
                errorAlert: false,
                perPage: 7,
                currentPage: 1,
                totalRows:0,
                formAlert:{
                    statusCode:0,
                    errorMsg: '',
                },
                formError: false
                }

            },
        methods: {
            getPlayers(){
                this.$http.get(this.$hostname+'/get_tournaments')
                    .then(response => {
                        this.tournaments = response.data.msg;
                        this.statusCode = response.data.code;
                        this.totalRows = this.tournaments.length;
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
            handleOK(evt){
                EventBus.$emit('submit',evt);
            },
            handleHide(){
                EventBus.$emit('show');
            },
            updTournaments(newTournament){
                this.tournaments.push(newTournament);
                this.$refs.tbl.refresh();
            },
            showDetails(details){
                this.$refs.td.title +=details.name;
                this.$refs.td.show();
                EventBus.$emit('init',details.id);
            },
            showFormAlert(formAlert){
                this.formError = true;
                this.formAlert = formAlert;
            }
        },
        created: function () {
            this.getPlayers();

        },
        mounted: function () {
            EventBus.$on('announce',this.updTournaments);
            EventBus.$on('formError',this.showFormAlert);

        }



    }
</script>

<style scoped src="../styles/tournament.css">

</style>