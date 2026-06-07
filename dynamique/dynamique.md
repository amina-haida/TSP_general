meilleur de celles visites + arriveer à=au nôtre 

liste_pays : ensemble {string}

// créer la structure bool où 1 signifie que le pays est visité

rec_tsp( visites ensemble,depart, arrivee string, memo_cout)
    if visites = {depart} :
        visites.append(arrive)
        memo[visites, arrive] = Distance(arrive, depart)
  
    else : 
        min_cout =+infini
        ville_min := string

        for pays dans visites :
            if pays != depart & memo_cout[visites privé de pays, pays] + 
            Distance(ville, arrivee) < min_cout :
                min_cout=memo_cout[visites privé de pays, pays] + Distance(pays, arrive)
                pays_min = pays
        
        memo_cout[visites , arrive] = min_cout

            
        

