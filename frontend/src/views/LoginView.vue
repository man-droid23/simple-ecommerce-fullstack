<template>
    <div class="box" v-bind:style="{width: '35vw', margin:'auto'}">
        <div class="content" v-bind:style="{width: '30vw'}">
            <h1>Login</h1>
            <hr>
            <form @submit.prevent="login">
                <div class="field">
                    <label class="label">Email</label>
                    <div class="control">
                        <input class="input" type="email" placeholder="Email" v-model="form.email">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Password</label>
                    <div class="control">
                        <input class="input" type="password" placeholder="Password" v-model="form.password">
                    </div>
                </div>
                <div class="field">
                    <div class="control">
                        <button class="button is-link">Login</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default{
    data(){
        return{
            form: {
                email: '',
                password: ''
            }
        }
    },
    methods: {
        async login(){
            console.log(this.form)
            try {
                const {data} = await axios.post('http://localhost:3000/auth/login', this.form)
                // set token to cookie
                document.cookie = `token=${data.token}`
                alert(data.message)
                this.$router.push('/')
            } catch (error) {
                alert(error)
                console.log(error)
            }
        }
    }
}
</script>