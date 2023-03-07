# poke-game-server

Exercício Typescript

# Ideia

Criar um jogo de cartas, estilo super trunfo. As regras básicas seguem listadas:

- As partidas serão compostas por no mínimo 2 e no máximo 8 jogadores
- Cada jogador receberá 9 cartas, aleatoriamente
- Será realizado um sorteio para determinar quem inicia o jogo
- O jogador vencedor da rodada anterior será o que inicia a próxima rodada
- A cada rodada o jogador da vez deverá escolher uma das caracteristicas listadas em uma de suas cartas. Os outros participantes escolhem uma de suas cartas para comparar a característica com a do jogador da vez. Vence o jogador que tiver a maior pontuação
  - Em caso de empate o jogador da vez deverá escolher outra característica, até que um dos participantes seja vencedor
- A carta utilizada é descartada do jogo
- O jogo termina quando todas as cartas forem utilizadas
- Ganha quem tiver vencido mais rodadas
  - Em caso de empate haverá a rodada de desempate. Os jogadores receberão uma nova carta e o vencedor da última rodada será o jogador da vez.
