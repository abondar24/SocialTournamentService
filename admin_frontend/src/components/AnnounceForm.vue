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
                formAlert:{
                statusCode:0,
                errorMsg: '',
                },
                tournamentId: null
            }
        },
        methods: {
            announceTournament() {
               if (this.tournament.name!=='' && this.tournament.deposit!==null){
                   this.$http.post(this.$hostname+'/announce_tournament?name='
                       +this.tournament.name+'&deposit='+this.tournament.deposit)
                       .then(response => {
                           this.tournamentId = response.data.msg;
                           this.formAlert.statusCode = response.data.code;
                           this.tournament.id = this.tournamentId;
                           if (this.formAlert.statusCode!==201){
                               this.formAlert.errorMsg = response.data;
                               this.formAlert.errorAlert = true;
                               EventBus.$emit('formError',this.formAlert);
                           } else {
                               EventBus.$emit('announce', this.tournament);
                           }
                       }, error => {
                           this.formAlert.errorMsg = error.statusText || error.data;
                           this.formAlert.statusCode = error.statusCode || error.status;
                           this.formAlert.errorAlert = true;
                           EventBus.$emit('formError',this.formAlert);
                       });
               }  else {
                   this.formAlert.errorAlert = true;
                   this.formAlert.errorMsg += 'Form filled incorrectly';
                   EventBus.$emit('formError',this.formAlert);
               }

            },
            resetForm() {
                Object.assign(this.$data, this.$options.data())
            }
        },
        created: function () {
            EventBus.$on('submit',this.announceTournament);
            EventBus.$on('show',this.resetForm);

        }

    }
</script>

<style scoped src="../styles/announceForm.css">

</style>