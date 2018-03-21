<template src="../templates/player.html">
    
</template>

<script>
    export default {
        name: "Player",

        data() {
            return {
                // fields: {
                //    id: {
                //         label: 'Id'
                //     },
                //     name:{
                //         label: 'Name'
                //     },
                //     points:  {
                //         label: 'Balance',
                //         sortable: true
                //     } },
                fields: [
                    {
                        key:'id',
                        label: 'Id'
                    },
                    {
                        key:'name',
                        label: 'Name'
                    },   {
                        key:'points',
                        label: 'Balance',
                        sortable: true
                    } ],

                players: [],
                statusCode: null,
                errorMsg: '',
                errorAlert: false,
            }
        },
        methods: {
            getPlayers(){
                this.$http.get(this.$hostname+'/get_players')
                    .then(response => {
                        this.players = response.data.msg;
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

<style scoped src="../styles/player.css">

</style>