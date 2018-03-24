<template src="../templates/announceForm.html">
    
</template>

<script>
    import EventBus from './event-bus';
    export default {
        name: "AnnounceForm",
        data() {
            return {

                tournament: {
                    id:null,
                    name:'',
                    deposit: null

                },
                statusCode: null,
                errorMsg: '',
                errorAlert: false,
                tournamentId: null
            }
        },
        methods: {
            announceTournament() {
               this.$http.post(this.$hostname+'/announce_tournament?name='
                   +this.tournament.name+'&deposit='+this.tournament.deposit)
            .then(response => {
                       this.tournamentId = response.data.msg;
                       this.statusCode = response.data.code;
                       this.tournament.id = this.tournamentId;
                       if (this.statusCode!==201){
                           this.errorMsg = response.data;
                           this.errorAlert = true;
                       } else {
                           EventBus.$emit('announce', this.tournament);
                       }
                   }, error => {
                       this.errorMsg = error.statusText || error.data;
                       this.statusCode = error.statusCode || error.status;
                       this.errorAlert = true;
                   });


            },
            resetForm() {
                this.tournament.name = '';
                this.tournament.deposit = null;
            }
        },
        created: function () {
            EventBus.$on('submit',this.announceTournament);
            EventBus.$on('reset',this.resetForm);
        }

    }
</script>

<style scoped src="../styles/announceForm.css">

</style>