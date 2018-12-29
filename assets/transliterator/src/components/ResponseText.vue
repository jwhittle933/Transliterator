<template>
    <div></div>
</template>

<script>
export default {
    name: "ResponseComponent",
    props: ["response"],
    data(){
        return {
            responseText: this.props.response
        }
    }
}
</script>

<style scoped>

</style>

