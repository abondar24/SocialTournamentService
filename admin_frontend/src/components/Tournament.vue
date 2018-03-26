<template src="../templates/tournament.html">
    
</template>

<script>
    import AnnounceForm from './AnnounceForm'
    import EventBus from './event-bus';
    export default {
        name: "Tournament",
        components:{
            AnnounceForm
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
            handleOK(){
                EventBus.$emit('submit');
            },
            handleCancel(){
                EventBus.$emit('reset');
            },
            updTournaments(newTournament){
                this.tournaments.push(newTournament);
            }
        },
        created: function () {
            this.getPlayers();

        },
        mounted: function () {
            EventBus.$on('announce',this.updTournaments);
        }



    }
</script>

<style scoped src="../styles/tournament.css">

</style>