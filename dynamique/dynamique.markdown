meilleur de celles visites + arriveer au nôtre 

liste_pays : ensemble {string}

faire structure Etat 
faire structure ensemble 
faire methode du nombre element present dans un ensemble 




cout_dynamique( ensemble_visites ensemble ,depart, arrive string, memo_cout [Etat]float)
    if ensemble_visites = {} :
        res := Distance(depart, arrive)
        etat = Etat{ depart , arrive}
        memo_cout[ etat] = res
        return res
  
    else : 
        min_cout =+infini
        ville_min := string
        etat = Etat {ensemble visites, arrivee} 
        for pays dans ensemble_visites :
            etat.ensemble = etat.ensemble.delete(pays)
            etat.arrivee = pays
            si etat dans memo_cout alors 
                res1 = memocout[Etat] + Distance( pays, arrivee)
            sinon 
                res1 = cout_dynamique( etat.ensemble, depart, etat.arrivee, memo_cout) + Distance( pays, arrivee)
            si res1 < min_cout alors 
                    min_cout = res1 
                    pays_min = pays 
            etat.ensemble = etat.ensemble.append(pays) 
            etat.arrivee = arrivee
            memo_cout[etat] = min_cout 
            return min_cout
                
            
        

