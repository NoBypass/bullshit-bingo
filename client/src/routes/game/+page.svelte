<script lang="ts">
    import { page } from "$app/stores";
    import Word from "$lib/components/Word.svelte";

    let words: string[] = [
      "this",
      "is",
      "a",
      "tesdsf asdjjklh a adsfadsf asdfas dt",
      "this",
      "is",
      "a",
      "test",
      "test",
      "this",
      "is",
      "a",
      "test",
      "this",
      "is",
      "a",
      "test",
      "this",
      "is",
      "a",
      "test",
      "this",
      "is",
      "a",
      "test",
    ];
  
</script>

<h1>UUID: {$page.url.searchParams.get("uuid")}</h1>
<section class="grid m-4 grid-cols-5 gap-4">
  {#each words as word}
    <Word>{word}</Word>
  {/each}
</section>
