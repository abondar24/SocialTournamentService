<template src="../templates/tournament.html">
    
</template>

<script>
    export default {
        name: "Tournament",
        data() {
            return {

                tournaments: [],
                statusCode: null,
                errorMsg: '',
                errorAlert: false,
            }
        },
        methods: {
            getPlayers(){
                this.$http.get(this.$hostname+'/get_tournaments')
                    .then(response => {
                        this.tournaments = response.data.msg;
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
            this.getPlayers();
        }
    }
</script>

<style scoped src="../styles/tournament.css">

</style>