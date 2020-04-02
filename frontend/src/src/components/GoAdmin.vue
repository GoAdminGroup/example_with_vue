<template>
    <div class="hello">
        <h1>{{ msg }}</h1>
        <ul>
            <li>
                <a href="https://doc.go-admin.cn" target="_blank">Docs</a>
            </li>
            <li>
                <a href="https://github.com/GoAdminGroup/go-admin" target="_blank">Github</a>
            </li>
            <li>
                <a href="https://www.go-admin.com" target="_blank">Official</a>
            </li>
        </ul>
    </div>
</template>

<script>
    export default {
        name: 'GoAdmin',
        data() {
            return {
                msg: 'GoAdmin'
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h1, h2 {
        font-weight: normal;
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
