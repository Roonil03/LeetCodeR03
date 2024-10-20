class Solution {
    public String intToRoman(int num) {
        String roman = "";
        int multiplier = 1;
        while(num!= 0)
        {
            if(num%10 == 4)
            {
                roman = charToReturn((num%10)*multiplier) + roman;
                num -= 4;
            }
            else if(num%10 == 9)
            {
                roman = charToReturn((num%10)*multiplier) + roman;
                num -= 9;
            }
            else if(num%10 != 0){
                if(num%10-5 >= 1)
                {
                    roman = charToReturn((1)*multiplier) + roman;
                    num -= 1;
                }
                else if(num%10>=5 && num%10 <10)
                {
                    roman = charToReturn((5)*multiplier) + roman;
                    num -= 5;
                }
                else{
                    roman = charToReturn((1)*multiplier) + roman;
                    num -= 1;
                }
            }
            //System.out.println(roman);
            if(num%10 == 0)
            {
                num /= 10;
                multiplier *= 10;
            }
        }
        return roman;
    }
    private String charToReturn(int n)
    {
        //String n = "";
        switch(n)
        {
            case 900:
            return "CM";
            case 400:
            return "CD";
            case 90:
            return "XC";
            case 40:
            return "XL";
            case 9:
            return "IX";
            case 4:
            return "IV";
            case 1:
            return "I";
            case 5:
            return "V";
            case 10:
            return "X";
            case 50:
            return "L";
            case 100:
            return "C";
            case 500:
            return "D";
            case 1000:
            return "M";
            default:
            return "ERROR";
        }
    }
}