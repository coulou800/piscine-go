cd the-final-cl-test/mystery

vari=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)

echo $vari

cat interviews/interview-"$vari"

echo $MAIN_SUSPECT
