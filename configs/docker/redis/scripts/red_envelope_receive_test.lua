local function  getRandomAmount(amount, num)
    if num == 1 or amount == 0 then
        return amount
    end

    local minAmount = 1
    local maxAmount = math.floor(amount/num)*2 - minAmount

    if maxAmount <= 0 then
        maxAmount = minAmount
    end

    return math.random(minAmount, maxAmount)
end

math.randomseed(os.time())

local total = 100
local count = 100
local num = count
local amount = 0
local dist = 0

for i=1,count do
    amount = getRandomAmount(total, num)
    num = num - 1
    dist = dist + amount
    total = total - amount

    print(total, num, dist, amount)
end