// Filename can be any, but package name must
// be the same with parent folder
package remote

func Hello() string {
    // Methods, structs, variables in the same package can be accessed freely among .go files
    return SayHi()
}