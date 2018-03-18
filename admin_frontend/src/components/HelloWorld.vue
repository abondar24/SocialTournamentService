<template>
  <div class="hello">
    <p>
      For guide and recipes on how to configure / customize this project,<br>
      check out the
      <a href="https://github.com/vuejs/vue-cli/tree/dev/docs" target="_blank">vue-cli documentation</a>.
    </p>
    <h1>{{ msg }}</h1>
    <h1>{{ statusCode }}</h1>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String,
    statusCode: String
  },
    methods: {
        status: function () {
            this.$http.get('http://localhost:8217/status')
                .then(response => {
                    console.log(response);
                this.msg = response.data.msg;
                this.statusCode = response.data.code;
        });
        }
    },
    created: function () {

        this.status();
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
