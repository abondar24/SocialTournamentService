<template src="../templates/player.html">
    
</template>

<script>
    import Balance from './Balance.vue'
    import EventBus from './event-bus';
    export default {
        name: "Player",
    components:{
            Balance
    },
        data() {
            return {
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
                perPage: 7,
                currentPage: 1,
                totalRows:0,
                balance: false,
                balanceError:false
            }
        },
        methods: {
            getPlayers(){
                this.$http.get(this.$hostname+'/get_players')
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
            openBalance(){
                this.balance = true;
            },
            handleBalanceError(err){
                this.balanceError = err.errorAlert;
                this.statusCode = err.statusCode;
                this.errorMsg = err.errorMsg;
            }
        },
        created: function () {
            this.getPlayers();

        },
        mounted: function () {
            EventBus.$on('balanceError',this.handleBalanceError);
        }


    }
</script>

<style scoped src="../styles/player.css">

</style>