<template>
    <div id="app">
        <Welcome />
    </div>
</template>

<script>
import Welcome from "./Welcome";
export default {
    name: "App",
    components: { Welcome }
};
</script>

<style>
#app {
    padding: 0;
    margin: 0;
    box-sizing: border-box;
}
</style>
