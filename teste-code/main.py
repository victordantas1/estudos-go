FORCA_RYGAR = 40
ENERGIA_RYGAR = 100

class Monstro:
	def __init(self, energia, forca):
		self.energia = energia
		self.forca = forca
	
	def __lt__(self, outro):
		if round(self.energia / FORCA_RYGAR * self.forca) > round(self.energia / FORCA_RAGNAR * self.forca):
			return True
		else: return False

def selection_sort(lis):
	for i in range(len(lis) - 1):
		menor = i
		for j in range(i + 1, len(lis)):
			if lis[menor] > lis[j]:
				menor = j
		lis[i], lis[menor] = lis[menor], lis[i]

def batalha(monstros, rygar):
	qtd_monstros = 0
	for monstro in monstros:
		while monstro.energia > 0 and rygar["energia"] > 0:
			monstro.energia -= rygar["forca"]
			rygar["energia"] -= monstro.forca
		if rygar["energia"] > 0:
			qtd_monstros += 1
	return qtd_monstros, rygar["energia"]
rygar = {"energia":100, "forca":40}
monstros = eval(input())
monstros = Monstro(monstro[0], monstro[1]) for monstro in montros
selection_sort(monstros)
qtd_monstros, energia_rygar = batalha(monstros, rygar)
print(qtd_monstros, energia_rygar)