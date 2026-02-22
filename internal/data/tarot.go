package data

type TarotCard struct {
	Name     string
	Arcana   string
	Upright  string
	Reversed string
}

var TarotDeck = []TarotCard{
	// Major Arcana
	{"The Fool", "Major", "Beginnings, innocence, spontaneity", "Recklessness, taken advantage of, inconsideration"},
	{"The Magician", "Major", "Power, skill, concentration", "Manipulation, poor planning, untapped talents"},
	{"The High Priestess", "Major", "Intuition, mystery, subconscious", "Secrets, disconnected from intuition, withdrawal"},
	{"The Empress", "Major", "Fertility, nurturing, abundance", "Dependence, smothering, emptiness"},
	{"The Emperor", "Major", "Authority, structure, control", "Tyranny, rigidity, coldness"},
	{"The Hierophant", "Major", "Tradition, conformity, morality", "Rebellion, subversiveness, new approaches"},
	{"The Lovers", "Major", "Love, harmony, relationships", "Imbalance, misalignment, disharmony"},
	{"The Chariot", "Major", "Direction, control, willpower", "Lack of control, opposition, no direction"},
	{"Strength", "Major", "Courage, patience, compassion", "Self-doubt, weakness, insecurity"},
	{"The Hermit", "Major", "Solitude, introspection, guidance", "Isolation, loneliness, withdrawal"},
	{"Wheel of Fortune", "Major", "Change, cycles, fate", "Bad luck, resistance to change, breaking cycles"},
	{"Justice", "Major", "Fairness, truth, law", "Unfairness, dishonesty, unaccountability"},
	{"The Hanged Man", "Major", "Surrender, letting go, new perspectives", "Delays, resistance, stalling"},
	{"Death", "Major", "Endings, transformation, transition", "Resistance to change, inability to move on"},
	{"Temperance", "Major", "Balance, moderation, patience", "Imbalance, excess, lack of long-term vision"},
	{"The Devil", "Major", "Bondage, addiction, materialism", "Release, restoring control, freedom"},
	{"The Tower", "Major", "Disaster, upheaval, sudden change", "Avoidance of disaster, fear of change"},
	{"The Star", "Major", "Hope, inspiration, generosity", "Lack of faith, despair, discouragement"},
	{"The Moon", "Major", "Illusion, fear, subconscious", "Release of fear, repressed emotion, clarity"},
	{"The Sun", "Major", "Joy, success, vitality", "Temporary depression, lack of success"},
	{"Judgement", "Major", "Reflection, reckoning, awakening", "Self-doubt, refusal of self-examination"},
	{"The World", "Major", "Completion, accomplishment, travel", "Incompletion, no closure, stagnation"},

	// Wands
	{"Ace of Wands", "Minor", "Creation, willpower, inspiration", "Lack of energy, lack of passion, boredom"},
	{"Two of Wands", "Minor", "Planning, making decisions, leaving home", "Fear of change, playing safe, bad planning"},
	{"Three of Wands", "Minor", "Looking ahead, expansion, rapid growth", "Obstacles, delays, frustration"},
	{"Four of Wands", "Minor", "Celebration, harmony, homecoming", "Conflict, transition, lack of support"},
	{"Five of Wands", "Minor", "Competition, rivalry, conflict", "Avoiding conflict, respecting differences"},
	{"Six of Wands", "Minor", "Victory, success, public reward", "Excess pride, lack of recognition, punishment"},
	{"Seven of Wands", "Minor", "Perseverance, defensive, maintaining control", "Give up, destroyed confidence, overwhelmed"},
	{"Eight of Wands", "Minor", "Rapid action, movement, quick decisions", "Panic, waiting, slowdown"},
	{"Nine of Wands", "Minor", "Resilience, grit, last stand", "Exhaustion, fatigue, questioning motivations"},
	{"Ten of Wands", "Minor", "Accomplishment, responsibility, burden", "Inability to delegate, overstressed, burnt out"},
	{"Page of Wands", "Minor", "Exploration, excitement, freedom", "Lack of direction, procrastination, creating conflict"},
	{"Knight of Wands", "Minor", "Action, adventure, fearlessness", "Anger, impulsiveness, recklessness"},
	{"Queen of Wands", "Minor", "Courage, confidence, independence", "Selfishness, jealousy, insecurities"},
	{"King of Wands", "Minor", "Big picture, leader, overcoming challenges", "Impulsive, overbearing, unachievable expectations"},

	// Cups
	{"Ace of Cups", "Minor", "New feelings, spirituality, intuition", "Emotional loss, blocked creativity, emptiness"},
	{"Two of Cups", "Minor", "Unity, partnership, connection", "Imbalance, broken communication, tension"},
	{"Three of Cups", "Minor", "Friendship, community, happiness", "Overindulgence, gossip, isolation"},
	{"Four of Cups", "Minor", "Apathy, contemplation, disconnectedness", "Sudden awareness, choosing happiness, acceptance"},
	{"Five of Cups", "Minor", "Loss, grief, self-pity", "Acceptance, moving on, finding peace"},
	{"Six of Cups", "Minor", "Familiarity, happy memories, healing", "Moving forward, leaving home, independence"},
	{"Seven of Cups", "Minor", "Choices, fantasy, illusion", "Lack of purpose, diversion, confusion"},
	{"Eight of Cups", "Minor", "Walking away, disillusionment, leaving behind", "Avoidance, fear of change, fear of loss"},
	{"Nine of Cups", "Minor", "Satisfaction, emotional stability, luxury", "Lack of inner joy, smugness, dissatisfaction"},
	{"Ten of Cups", "Minor", "Inner happiness, fulfillment, dreams coming true", "Shattered dreams, broken family, domestic disharmony"},
	{"Page of Cups", "Minor", "Happy surprise, dreamer, sensitivity", "Emotional immaturity, insecurity, disappointment"},
	{"Knight of Cups", "Minor", "Following the heart, idealist, romantic", "Moodiness, disappointment, unrealistic"},
	{"Queen of Cups", "Minor", "Compassion, calm, comfort", "Martyrdom, insecurity, dependence"},
	{"King of Cups", "Minor", "Emotional balance, diplomacy, generosity", "Coldness, moodiness, bad advice"},

	// Swords
	{"Ace of Swords", "Minor", "Breakthrough, clarity, sharp mind", "Confusion, brutality, chaos"},
	{"Two of Swords", "Minor", "Difficult choices, indecision, stalemate", "Lesser of two evils, no right choice, confusion"},
	{"Three of Swords", "Minor", "Heartbreak, suffering, grief", "Recovery, forgiveness, moving on"},
	{"Four of Swords", "Minor", "Rest, restoration, contemplation", "Restlessness, burnout, stress"},
	{"Five of Swords", "Minor", "Unbridled ambition, win at all costs, sneakiness", "Lingering resentment, desire to reconcile, forgiveness"},
	{"Six of Swords", "Minor", "Transition, leaving behind, moving on", "Emotional baggage, unresolved issues, resisting transition"},
	{"Seven of Swords", "Minor", "Deception, trickery, tactics", "Imposter syndrome, self-deceit, keeping secrets"},
	{"Eight of Swords", "Minor", "Imprisonment, entrapment, self-victimization", "Self-acceptance, new perspective, freedom"},
	{"Nine of Swords", "Minor", "Anxiety, hopelessness, trauma", "Hope, reaching out, despair"},
	{"Ten of Swords", "Minor", "Failure, collapse, defeat", "Recovery, regeneration, fear of ruin"},
	{"Page of Swords", "Minor", "Curiosity, restlessness, mental energy", "Deception, manipulation, all talk"},
	{"Knight of Swords", "Minor", "Action, impulsiveness, defending beliefs", "No direction, disregard for consequences, unpredictability"},
	{"Queen of Swords", "Minor", "Complexity, perceptiveness, clear mindedness", "Cold-hearted, cruel, bitterness"},
	{"King of Swords", "Minor", "Head over heart, discipline, truth", "Manipulative, cruel, weakness"},

	// Pentacles
	{"Ace of Pentacles", "Minor", "Opportunity, prosperity, new venture", "Lost opportunity, missed chance, bad investment"},
	{"Two of Pentacles", "Minor", "Balancing decisions, priorities, adapting to change", "Loss of balance, disorganized, overwhelmed"},
	{"Three of Pentacles", "Minor", "Teamwork, collaboration, learning", "Disharmony, misalignment, working alone"},
	{"Four of Pentacles", "Minor", "Conservation, frugality, security", "Greediness, stinginess, possessiveness"},
	{"Five of Pentacles", "Minor", "Financial loss, poverty, lack mindset", "Recovery, charity, improvement"},
	{"Six of Pentacles", "Minor", "Charity, generosity, sharing", "Strings attached, stinginess, power and domination"},
	{"Seven of Pentacles", "Minor", "Vision, perseverance, profit", "Lack of vision, limited reward, no growth"},
	{"Eight of Pentacles", "Minor", "Apprenticeship, passion, high standards", "Lack of passion, uninspired, no motivation"},
	{"Nine of Pentacles", "Minor", "Fruits of labor, rewards, luxury", "Reckless spending, living beyond means, false success"},
	{"Ten of Pentacles", "Minor", "Wealth, inheritance, family", "Financial failure, lone wolf, loss of tradition"},
	{"Page of Pentacles", "Minor", "Ambition, desire, diligence", "Lack of commitment, greediness, laziness"},
	{"Knight of Pentacles", "Minor", "Efficiency, hard work, responsibility", "Laziness, obsessiveness, work without reward"},
	{"Queen of Pentacles", "Minor", "Practicality, creature comforts, financial security", "Self-centeredness, jealousy, smothering"},
	{"King of Pentacles", "Minor", "Wealth, business, leadership", "Financially inept, obsessed with wealth, stubborn"},
}
