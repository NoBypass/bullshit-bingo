<script lang="ts">
    import { goto } from "$app/navigation";
    import Button from "$lib/components/Button.svelte";

    let uuid: string = "";

</script>
<section class="flex h-full w-full flex-col gap-4 p-4">
  <Button on:click={() => goto("/game")}>New Game</Button>
  <div class="w-full flex flex-row gap-4">
    <input type="text" bind:value={uuid} class="border-2 p-1 border-black rounded justify-between grow " /> 
    <Button on:click={() => goto(`/game?uuid=${uuid}`)}>join</Button>
  </div>
  
</section>
