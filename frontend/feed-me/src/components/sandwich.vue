<template>
    <div id="sandwich">
        <h2>Feed me more</h2>
        <form v-if="!submitted">
            <label>sandwich id:</label>
            <input type="number" v-model.lazy="sandwich.id" required />
            <label>sandwich name:</label>
            <textarea v-model.lazy.trim="sandwich.name"></textarea>
            <button v-on:click.prevent="post">Add sandwich</button>
            <hr>
            <button v-on:click.prevent="put(sandwich.id)">Update sandwich</button>
            <button v-on:click.prevent="del(sandwich.id)">delete sandwich</button>
            <button v-on:click.prevent="get">Get sandwichs</button>
        </form>
        <div v-if="submitted">
            <h3>Thanks for adding your sandwich</h3>
        </div>
        <div v-for="sandwich in sandwichs" :key="sandwich.id"  id="preview">
            
            <h3>Preview sandwich</h3>
            <p>sandwich id: {{ sandwich.id }}</p>
            <p>sandwich name: {{ sandwich.name }}</p>
        </div>
    </div>
</template>

<script>
// Imports

export default {
    data () {
        return {
            sandwichs:[],
            sandwich: {
                id: '',
                name: '',
            },
            submitted: false
        }
    },
    methods: {
        post: function(){
            this.$http.post('http://localhost:8000/api/sandwiches', {
                id: parseInt(this.sandwich.id),
                name: this.sandwich.name,
            }).then(function(){

                this.submitted = false;
            });
        },
        get: function(){
            this.$http.get('http://localhost:8000/api/sandwiches', {
            }).then(function(data){
                if(data.status == "200"){ console.log(data); }
                // sandwichs.append(data)
                this.sandwichs=data.body
                console.log(this.sandwichs)
                this.submitted = false;
            });
        },
        put: function(id){
            this.$http.put(`http://localhost:8000/api/sandwiches/` + id, {
                id: parseInt(id),
                name: this.sandwich.name,
            }).then(function(){
                this.submitted = false;
            });
        },
        del: function(id){
            this.$http.delete(`http://localhost:8000/api/sandwiches/` + id, {
                id: parseInt(id),
            }).then(function(){
                this.submitted = false;
            });
        }
    }
}
</script>

<style>
#sandwich *{
    box-sizing: border-box;
}
#sandwich{
    margin: 20px auto;
    max-width: 500px;
}
label{
    display: block;
    margin: 20px 0 10px;
}
input[type="text"], textarea{
    display: block;
    width: 100%;
    padding: 8px;
}
#preview{
    padding: 10px 20px;
    border: 1px dotted #ccc;
    margin: 30px 0;
}
h3{
    margin-top: 10px;
}
#checkboxes input{
    display: inline-block;
    margin-right: 10px;
}
#checkboxes label{
    display: inline-block;
    margin-top: 0;
}
</style>
