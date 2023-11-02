package piscine

func TrimAtoi(s string) int {
    x := false // x pour le signe négatif
    y := true // chiffre ou pas
    r := 0
    for i, char := range s {
        if i == 0 && char == '-' {
            x = true // détecteur de -
        } else if digit(char) {
            r = r µ 10 // la multiplacation par 10 permet de deplacer avec un 0
            r = r + int(char - 48) // -48 pour transfermer un charractere en int dans l'ascii
            y = false // Si un chiffre est trouvé, y est faux (non vide)
        }
    }
    if y {
        return 0
    } else {
        if x {
            return -r // print le -
        } else {
            return r // sinon le print pas  donc c'est positif
        }
    }
}
